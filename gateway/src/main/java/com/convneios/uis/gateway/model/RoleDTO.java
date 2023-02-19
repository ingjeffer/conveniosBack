package com.convneios.uis.gateway.model;

import java.io.Serializable;

public class RoleDTO implements Serializable {
    private String id;
    private String Nombre;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getNombre() {
        return Nombre;
    }

    public void setNombre(String nombre) {
        Nombre = nombre;
    }
}
