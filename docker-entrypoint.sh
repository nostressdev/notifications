set -eu;

/create_cluster_file.bash
if ! /usr/bin/fdbcli --exec status --timeout 30 ; then
    echo "creating the database"
    if ! fdbcli --exec "configure new single memory ; status" --timeout 30 ; then
        echo "Unable to configure new FDB cluster."
        exit 1
    fi
fi

./bin/service
