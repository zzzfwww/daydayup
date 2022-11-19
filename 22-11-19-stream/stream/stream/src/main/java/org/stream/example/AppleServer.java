package org.stream.example;

import org.stream.example.entity.Apple;
import org.testng.annotations.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Predicate;
import java.util.stream.Collectors;

public class AppleServer {
    private static List<Apple> appleStore = new ArrayList<>();

    static {
        appleStore.add(new Apple(1, "red", 500, "hunan"));
        appleStore.add(new Apple(2, "red", 400, "hunan"));
        appleStore.add(new Apple(3, "green", 300, "hunan"));
        appleStore.add(new Apple(4, "green", 200, "guangdong"));
        appleStore.add(new Apple(5, "green", 100, "guangdong"));
    }

    public static void main(String[] args) {
        System.out.println("Hello world!");
        test2();
        new AppleServer().testFun(a -> a.getColor().equals("red") && a.getWeight() > 400);
        testGroupStream();
    }

    // steam 查找
    @Test
    public static void test2() {
        List<Apple> red = appleStore.stream()
                .filter(a -> a.getColor().equals("red"))
                // 增加重量判断
                .filter(a -> a.getWeight() > 400)
                .collect(Collectors.toList());
        System.out.println(red);
    }

    public static void testGroupStream() {
        appleStore.stream().collect(Collectors.groupingBy(a -> a.getColor(),
                Collectors.averagingInt(a -> a.getWeight()))).forEach((k, v) -> System.out.println(k + ":" + v));
    }

    public void testFun(Predicate<? super Apple> pr) {
        List<Apple> collect = appleStore.stream()
                .filter(pr)
                .collect(Collectors.toList());
        System.out.println(collect);
    }

    // 找出红色的苹果
    public void test1() {
        for (Apple apple : appleStore) {
            if (apple.getColor().equals("red")) {
                // add
            }
        }
    }

    // 求出每个颜色的平均值
    public void testGroup() {
        // 1. 基于颜色分组
        Map<String, List<Apple>> maps = new HashMap<>();
        for (Apple apple : appleStore) {
            List<Apple> list = maps.computeIfAbsent(apple.getColor(), key -> new ArrayList<>());
            list.add(apple);
        }
        for (Map.Entry<String, List<Apple>> entry : maps.entrySet()) {
            int weights = 0;
            for (Apple apple : entry.getValue()) {
                weights += apple.getWeight();
            }
            System.out.println(String.format("颜色%s 平均重量：%s", entry.getKey(), weights / entry.getValue().size()));
        }
    }
}