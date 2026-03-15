package com.blog.util;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;

import java.util.Collection;
import java.util.Map;
import java.util.concurrent.TimeUnit;

@Component
public class RedisCache {

    @Autowired
    private RedisTemplate<String, Object> redisTemplate;

    public void setCacheObject(String key, Object value) {
        redisTemplate.opsForValue().set(key, value);
    }

    public void setCacheObject(String key, Object value, long timeout, TimeUnit timeUnit) {
        redisTemplate.opsForValue().set(key, value, timeout, timeUnit);
    }

    public <T> T getCacheObject(String key) {
        return (T) redisTemplate.opsForValue().get(key);
    }

    public void deleteObject(String key) {
        redisTemplate.delete(key);
    }

    public void deleteObject(Collection<String> keys) {
        redisTemplate.delete(keys);
    }

    public void setCacheMap(String key, Map<String, Object> map) {
        redisTemplate.opsForHash().putAll(key, map);
    }

    public <T> T getCacheMapValue(String key, String hashKey) {
        return (T) redisTemplate.opsForHash().get(key, hashKey);
    }

    public void setCacheMapValue(String key, String hashKey, Object value) {
        redisTemplate.opsForHash().put(key, hashKey, value);
    }

    public void increaseCacheMapValue(String key, String hashKey, long delta) {
        redisTemplate.opsForHash().increment(key, hashKey, delta);
    }

    public Map<Object, Object> getCacheMap(String key) {
        return redisTemplate.opsForHash().entries(key);
    }

    public void deleteCacheMapValue(String key, String hashKey) {
        redisTemplate.opsForHash().delete(key, hashKey);
    }

    public boolean hasKey(String key) {
        return Boolean.TRUE.equals(redisTemplate.hasKey(key));
    }

    public boolean expire(String key, long timeout, TimeUnit timeUnit) {
        return Boolean.TRUE.equals(redisTemplate.expire(key, timeout, timeUnit));
    }
}
