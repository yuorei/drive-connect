# ソースコードの出力先
PROTO_DEST=./proto
# protoファイルが置いてあるディレクトリ
PROTO_SRC=../proto

# Protocol BuffersからnodeのgRPCコードを自動生成
grpc_tools_node_protoc \
  --js_out=import_style=commonjs,binary:${PROTO_DEST} \
  --grpc_out=${PROTO_DEST} \
  --plugin=protoc-gen-grpc=$(which grpc_tools_node_protoc_plugin) \
  -I ${PROTO_SRC} \
  ${PROTO_SRC}/*

# typescript用の型定義を作成
grpc_tools_node_protoc \
  --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
  --ts_out=${PROTO_DEST} \
  -I ${PROTO_SRC} \
  ${PROTO_SRC}/*