if [ ! -d ../outpost/outpostrpc ]; then
    mkdir ../outpost/outpostrpc
fi

pb_prefix="text_item"

protoc --go_out=../outpost/outpostrpc --go_opt=paths=source_relative \
    --go-grpc_out=../outpost/outpostrpc --go-grpc_opt=paths=source_relative \
    ${pb_prefix}.proto

python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ${pb_prefix}.proto

if [ ! -f ../posten/env.mojo ]; then
    echo "var PY_PROTO_PATH: StringLiteral = \"$(pwd)\"" >> ../posten/env.mojo
    echo "var PB_2_MODULE: StringLiteral = \"${pb_prefix}_pb2\"" >> ../posten/env.mojo
    echo "var PB_2_GRPC_MODULE: StringLiteral = \"${pb_prefix}_pb2_grpc\"" >> ../posten/env.mojo
fi