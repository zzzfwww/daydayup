package com.example.demo.aop;

public class SaveOrderConvert implements Convert<SaveOrder> {

    @Override
    public OperateLogDO convert(SaveOrder saveOrder) {
        OperateLogDO operateLogDO = new OperateLogDO();
        operateLogDO.setOrderId(saveOrder.getId());
        return operateLogDO;
    }
}
