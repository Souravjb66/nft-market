package com.nft.market.project.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "Users")
public class User {
    private String name;
    @Id
    private String email;
    private int age;
    private int totalWallet;
    private String password;
    public User(){}
    public User(String name,String email,int age,int totalWallet,String password){
        this.age=age;
        this.email=email;
        this.name=name;
        this.password=password;
        this.totalWallet=totalWallet;
    }
    public String getName(){return name;}
    public String getEmail(){return email;}
    public int getAge(){return age;}
    public int getTotalWallet(){return totalWallet;}
    public String getPassword(){return password;}
    public void setName(String name){
        this.name=name;
    }
    public void setEmail(String email){
        this.email=email;
    }
    public void setAge(int age){
        this.age=age;
    }
    public void setPassword(String password){
        this.password=password;
    }
    public void setTotalWallet(int no){
        this.totalWallet=no;
    }
}
