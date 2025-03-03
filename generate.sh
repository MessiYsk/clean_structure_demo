# 生成接口代码
kitex -module github.com/MessiYsk/clean_structure_demo idl/repayment.thrift

# 接口测试代码
# kitexcall -idl-path idl/repayment.thrift -m ManualRepay -d '{"CreditCardID":"card123","Amount":1000.0,"Fee":10.0}' -e localhost:8888
