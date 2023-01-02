package com.convneios.uis.gateway.model;

import java.io.Serializable;

public class SessionDTO implements Serializable {

    private static final long serialVersionUID = 1L;
    private String id;
    private long exp;


    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public long getExp() {
        return exp;
    }

    public void setExp(long exp) {
        this.exp = exp;
    }
}
