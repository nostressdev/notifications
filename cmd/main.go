package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nostressdev/notifications/internal/notifications"
	"github.com/nostressdev/notifications/internal/storage"
	"github.com/nostressdev/signer"

	"github.com/nostressdev/notifications/pushkit/push/config"
	pushkit "github.com/nostressdev/notifications/pushkit/push/core"
	"github.com/nostressdev/notifications/utils"

	pb "github.com/nostressdev/notifications/proto"

	"github.com/nostressdev/runner"

	firebase "firebase.google.com/go/v4"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const (
	authUrl = "https://login.cloud.huawei.com/oauth2/v2/token"
	pushUrl = "https://api.push.hicloud.com"
)

var (
	host      string
	port      string
	secretKey string

	// for firebase cloud messaging
	// service account's credentials: json string / path of json file
	firebsaseCreds string

	// for huawei pushkit
	appId     string
	appSecret string

	// for email sending
	smtpHost     string
	smtpPort     string
	smtpLogin    string
	smtpPassword string
	smtpSender   string
)

func loadEnvironmentVariables() {
	if host = os.Getenv("HOST"); host == "" {
		log.Fatalln("$HOST environment variable should be specified")
	}
	if port = os.Getenv("PORT"); port == "" {
		log.Fatalln("$PORT environment variable should be specified")
	}
	if secretKey = os.Getenv("SECRET_KEY"); secretKey == "" {
		log.Fatalln("$SECRET_KEY environment variable should be specified")
	}

	if firebsaseCreds = os.Getenv("FIREBASE_CREDS"); firebsaseCreds == "" {
		log.Fatalln("$FIREBASE_CREDS environment variable should be specified")
	}
	if appId = os.Getenv("APP_ID"); appId == "" {
		log.Fatalln("$APP_ID environment variable should be specified")
	}
	if appSecret = os.Getenv("APP_SECRET"); appSecret == "" {
		log.Fatalln("$APP_SECRET environment variable should be specified")
	}
	if smtpHost = os.Getenv("SMTP_HOST"); smtpHost == "" {
		log.Fatalln("SMTP_HOST environment variable should be specified")
	}
	if smtpPort = os.Getenv("SMTP_PORT"); smtpPort == "" {
		log.Fatalln("SMTP_PORT environment variable should be specified")
	}
	if smtpLogin = os.Getenv("SMTP_LOGIN"); smtpLogin == "" {
		log.Fatalln("SMTP_LOGIN environment variable should be specified")
	}
	if smtpPassword = os.Getenv("SMTP_PASSWORD"); smtpPassword == "" {
		log.Fatalln("SMTP_PASSWORD environment variable should be specified")
	}
	if smtpSender = os.Getenv("SMTP_SENDER"); smtpSender == "" {
		log.Fatalln("SMTP_SENDER environment variable should be specified")
	}
}

func createNetworkListener() net.Listener {
	addr := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving with addr: %v\n", addr)
	return listener
}

func createServiceConnection(serviceName string) *grpc.ClientConn {
	connURL := os.Getenv(serviceName + "_URL")
	if connURL == "" {
		log.Fatalln("$" + serviceName + "_URL environment variable should be specified")
	}
	conn, err := grpc.Dial(connURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect to " + serviceName)
	}
	return conn
}

func createConfigSMTP() *utils.SMTP {
	port, err := strconv.Atoi(strings.TrimSpace(smtpPort))
	if err != nil {
		log.Fatalf("cannot convert smtp port to int, \"%s\" : %s\n", smtpPort, err.Error())
	}
	return &utils.SMTP{
		Host:     smtpHost,
		Port:     port,
		Login:    smtpLogin,
		Sender:   smtpSender,
		Password: smtpPassword,
	}
}

type ServerResource struct {
	Server *grpc.Server
}

func (s *ServerResource) Init(ctx context.Context) error {
	fdb.MustAPIVersion(600)

	db := fdb.MustOpenDefault()

	repo := storage.NewNotificationsFDB(&storage.ConfigNotificationsFDB{
		DB:       db,
		Subspace: subspace.Sub("notifications_subspace"),
	})

	loadEnvironmentVariables()

	signer := signer.NewSignerJWT(signer.TokenProviderConfig{
		Expiration: time.Hour,
		SecretKey:  []byte(secretKey),
	})

	opts := option.WithCredentialsJSON([]byte(firebsaseCreds))
	app, err := firebase.NewApp(context.Background(), nil, opts)

	if err != nil {
		return fmt.Errorf("Could not initialize firebase app: %s", err.Error())
	}

	firebaseApp := &utils.FirebaseApp{
		App: app,
	}

	pushkitApp, err := pushkit.NewHttpClient(&config.Config{
		AppId:     appId,
		AppSecret: appSecret,
		AuthUrl:   authUrl,
		PushUrl:   pushUrl,
	})

	if err != nil {
		return fmt.Errorf("Could not initialize huawei pushkit app: %s", err.Error())
	}

	huaweiApp := &utils.HuaweiApp{
		App: pushkitApp,
	}

	s.Server = grpc.NewServer()
	pb.RegisterNotificationsServer(s.Server, notifications.New(&notifications.Config{
		Signer:      signer,
		Storage:     repo,
		FirebaseApp: firebaseApp,
		HuaweiApp:   huaweiApp,
		EmailApp: &utils.EmailApp{
			SMTP: createConfigSMTP(),
		},
	}))
	return nil
}

func (s *ServerResource) Release(ctx context.Context) error {
	s.Server.GracefulStop()
	return nil
}

type ListenerResource struct {
	Listener net.Listener
}

func (r *ListenerResource) Init(ctx context.Context) error {
	r.Listener = createNetworkListener()
	return nil
}

func (r *ListenerResource) Release(ctx context.Context) error {
	return nil
}

type MainJob struct {
	Server   *ServerResource
	Listener *ListenerResource
}

func (main *MainJob) Run() error {
	if err := main.Server.Server.Serve(main.Listener.Listener); err != nil {
		return err
	}
	return nil
}

func (main *MainJob) Shutdown(ctx context.Context) error {
	return nil
}

func main() {
	server := &ServerResource{}
	listener := &ListenerResource{}
	app := runner.New(runner.DefaultConfig(),
		[]runner.Resource{server, listener},
		[]runner.Job{&MainJob{
			Server:   server,
			Listener: listener,
		}})
	app.Run()
}
