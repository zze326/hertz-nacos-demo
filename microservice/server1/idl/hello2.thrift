namespace go hello2.example

struct HelloReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}

struct TestPostReq {
    // 名字
    1: string name (api.body="name");
}

struct TestPostResp {
    1: string Msg;
}


service Hello2Service {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello2");
    TestPostResp TestPost(1: TestPostReq request) (api.post="/test-post")
}