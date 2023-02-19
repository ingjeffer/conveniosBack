package com.convneios.uis.gateway.filter;

import com.auth0.jwt.JWT;
import com.convneios.uis.gateway.model.SessionDTO;
import com.convneios.uis.gateway.repository.SessionRepository;
import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.Base64;
import java.util.logging.Logger;

@Component
public class OAuthFilter implements GlobalFilter {

    @Autowired
    private SessionRepository sessionRedisRepository;

    private static final String SESSION_PATH = "/api/usuario/session";

    static Logger logger = Logger.getLogger(OAuthFilter.class.getName());

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {

        if (SESSION_PATH.equals(exchange.getRequest().getPath().toString())) {
            return chain.filter(exchange);
        }

        try {
            String[] splitToken = exchange.getRequest().getHeaders().get("Authorization").get(0).split("Bearer ");

            if (splitToken.length != 2) {
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }

            String decodedString = getTokenDecode(splitToken[1]);
            logger.info("jwt payload -> " + decodedString);

            Gson gson = new Gson();
            SessionDTO session = gson.fromJson(decodedString, SessionDTO.class);

            String token = sessionRedisRepository.getSession(session.getId());

            if (token == null || token.isEmpty()){
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }

            if (!token.equals(splitToken[1])){
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }

            if (session.getExp() < System.currentTimeMillis() / 1000L) {
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }
            exchange.getRequest()
                    .mutate()
                    .header("x-role", session.getRole().getNombre())
                    .build();

            exchange.getRequest()
                    .mutate()
                    .header("x-id", session.getId())
                    .build();

            return chain.filter(exchange);

        }catch (Exception e) {
            exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
            return exchange.getResponse().setComplete();
        }

    }

    private String getTokenDecode(String token1) {
        return new String(Base64.getDecoder().decode(JWT.decode(token1).getPayload()));
    }


}
