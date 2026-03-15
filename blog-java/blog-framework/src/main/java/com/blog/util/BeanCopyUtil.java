package com.blog.util;

import org.springframework.beans.BeanUtils;

import java.util.List;
import java.util.stream.Collectors;

public final class BeanCopyUtil {

    private BeanCopyUtil() {}

    public static <T> T copyBean(Object source, Class<T> target) {
        T result;
        try {
            result = target.getDeclaredConstructor().newInstance();
            BeanUtils.copyProperties(source, result);
        } catch (Exception e) {
            throw new RuntimeException("Bean copy failed", e);
        }
        return result;
    }

    public static <T> List<T> copyBeanList(List<?> sourceList, Class<T> target) {
        return sourceList.stream()
                .map(source -> copyBean(source, target))
                .collect(Collectors.toList());
    }
}
