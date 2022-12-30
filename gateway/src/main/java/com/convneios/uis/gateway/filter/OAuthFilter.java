package com.convneios.uis.gateway.filter;

import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.logging.Logger;

@Component
public class OAuthFilter implements GlobalFilter {

    private static final String SESSION_PATH = "/api/usuario/session";

    static Logger logger = Logger.getLogger(OAuthFilter.class.getName());

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {

        if (SESSION_PATH.equals(exchange.getRequest().getPath().toString())) {
            logger.info("ignorando al filtro");
            return chain.filter(exchange);
        }

        logger.info("entrando al filtro");
        return chain.filter(exchange);
    }
}
