# example: make new-microservice service_name=server1
new-microservice:
	hz new --out_dir /Users/zhangzhongen/work/git_repo/prj/hertz-demo/microservice/$(service_name)/ --mod hertz-demo/microservice/$(service_name) --idl /Users/zhangzhongen/work/git_repo/prj/hertz-demo/microservice/$(service_name)/idl/*.thrift