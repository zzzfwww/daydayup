package com.example.demo.aop;

import lombok.Data;

@Data
public class OperateLogDO {
    private Long orderId;
    private String desc;
    private String result;
}
