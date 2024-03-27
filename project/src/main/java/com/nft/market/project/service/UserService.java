package com.nft.market.project.service;

import com.nft.market.project.model.User;
import com.nft.market.project.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class UserService {
    @Autowired
    private UserRepository repository;

    public String saveUser(User user){
        try{
            this.repository.save(user);
            return "saved";
        }
        catch (Exception ex){
            System.out.println(ex);
            return "not saved"+ex;

        }

    }
    public User getUserFromPrimary(String Email){

        var value=this.repository.findById(Email);
        var user=value.stream().iterator().next();
        return user;

    }
    public List<User>getUserFromName(String name){
        var v=this.repository.findByname(name);
        return v;
    }
    public List<User>getUserFromTotalWallet(int i){
        var u=this.repository.findBytotalWallet(i);
        return u;
    }
    public List<User>getUserFromAge(int age){
        var u=this.repository.findByage(age);
        return u;
    }
    public void incrementWallet(String email){
        User u=getUserFromPrimary(email);
        var total=u.getTotalWallet();
        total=total+1;
        u.setTotalWallet(total);
        this.repository.save(u);

    }

}
