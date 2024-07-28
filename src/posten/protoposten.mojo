from python import Python
from env import PY_PROTO_PATH, PB_2_MODULE, PB_2_GRPC_MODULE

@value
struct TextItem:
    var text: String
    var id: String
    var timestamp: String
    var category: String


struct TextItemStubbe:
    var stub: PythonObject
    var text_item_pb: PythonObject

    fn __init__(inout self, channel_host: String) raises:
        var grpc = Python.import_module("grpc")
        Python.add_to_path(PY_PROTO_PATH)
        self.text_item_pb = Python.import_module(PB_2_MODULE)
        var text_item_pb2_grpc = Python.import_module(PB_2_GRPC_MODULE)
        var channel = grpc.insecure_channel(channel_host)
        self.stub = text_item_pb2_grpc.OutpostServiceStub(channel)

    fn retrieve(self) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveReq()
        var res = self.stub.Retrieve(req)
        print(res)
        return List[TextItem]()

    fn retrieve_on_id(self) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveOnIdReq(
            id=""
        )
        var res = self.stub.RetrieveOnId(req)
        print(res)
        return List[TextItem]()

    fn retrieve_on_time(self) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveOnTimeReq(
            start="",
            end=""
        )
        var res = self.stub.RetrieveOnTime(req)
        print(res)
        return List[TextItem]()

    fn retrieve_on_category(self) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveOnCategoryReq(
            category=""
        )
        var res = self.stub.RetrieveOnCategory(req)
        print(res)
        return List[TextItem]()


fn main() raises:
    var stub = TextItemStubbe("localhost:5000")
    var x = stub.retrieve()
    x = stub.retrieve_on_id()
    x = stub.retrieve_on_time()
    x = stub.retrieve_on_category()