package com.convneios.uis.gateway.repository.impl;

import com.convneios.uis.gateway.model.SessionDTO;
import com.convneios.uis.gateway.repository.SessionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;

@Component
public class SessionRedisRepository implements SessionRepository {

    @Autowired
    private RedisTemplate<String, SessionDTO> redisTemplate;


    public String getSession(String id){
        return String.valueOf(this.redisTemplate.opsForValue().get("Id"+id));
    }
}
