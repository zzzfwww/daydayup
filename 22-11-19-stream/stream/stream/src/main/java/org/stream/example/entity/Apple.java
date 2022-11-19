package org.stream.example.entity;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class Apple {
    private Integer id;
    private String color;
    private Integer weight;
    private String origin;
}
