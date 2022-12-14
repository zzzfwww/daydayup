package com.example.demo;

import com.example.demo.aop.OrderService;
import com.example.demo.aop.SaveOrder;
import com.example.demo.aop.UpdateOrder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application implements CommandLineRunner {
    public static void main(String[] args) {
        new SpringApplication(Application.class).run(args);
    }

    @Autowired
    private OrderService orderService;

    @Override
    public void run(String... args) throws Exception {
        SaveOrder saveOrder = new SaveOrder();
        saveOrder.setId(1L);

        orderService.saveOrder(saveOrder);

        UpdateOrder updateOrder = new UpdateOrder();
        updateOrder.setOrderId(100L);
        orderService.updateOrder(updateOrder);

    }
}
