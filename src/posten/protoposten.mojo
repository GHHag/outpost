from python import Python
from env import PY_PROTO_PATH, PB_2_MODULE, PB_2_GRPC_MODULE

@value
struct TextItem:
    var text: String
    var ref_tag: String
    var timestamp: String
    var category: String

    @staticmethod
    fn convert_repeated_text_items(
        owned text_items: PythonObject
    ) raises -> List[TextItem]:
        var converted_text_items = List[TextItem]()

        for i in range(len(text_items)):
            converted_text_items.append(
                TextItem(
                    text_items[i].text,
                    text_items[i].ref_tag,
                    text_items[i].timestamp,
                    text_items[i].category,
                )
            )

        return converted_text_items

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

    # TODO: insert method not needed in mojo service
    fn insert(self, text_item: TextItem) raises:
        var req = self.text_item_pb.TextItem(
            text=text_item.text,
            ref_tag=text_item.ref_tag,
            timestamp=text_item.timestamp,
            category=text_item.category
        )
        var res = self.stub.InsertTextItem(req)
        print("insert response:", res)

    fn retrieve(self) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveReq()
        var res = self.stub.Retrieve(req)
        var text_items = TextItem.convert_repeated_text_items(res.text_items)
        return text_items

    fn retrieve_on_ref_tag(self, ref_tag: String) raises -> List[TextItem]:
        var req = self.text_item_pb.RetrieveOnRefTagReq(ref_tag=ref_tag)
        var res = self.stub.RetrieveOnRefTag(req)
        var text_items = TextItem.convert_repeated_text_items(res.text_items)
        return text_items

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
    var stub = TextItemStubbe("localhost:5055")

    # var text_item = TextItem("test", "test", "test", "test")
    # stub.insert(text_item)
    # text_item = TextItem("test", "test", "test", "test")
    # stub.insert(text_item)

    # var x = stub.retrieve()
    # for i in range(len(x)):
    #     print(x[i].text, x[i].ref_tag, x[i].timestamp, x[i].category)

    var x = stub.retrieve_on_ref_tag("test")
    for i in range(len(x)):
        print(x[i].text, x[i].ref_tag, x[i].timestamp, x[i].category)
