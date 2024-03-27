package com.nft.market.project.controller;

import com.nft.market.project.model.User;
import com.nft.market.project.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class FirstController {
    @Autowired
    private UserService service;
    @GetMapping("/getuser-email/{email}")
    public User getUserByEmail(@PathVariable String email){
        var value=this.service.getUserFromPrimary(email);
        return value;

    }
    @PostMapping("/save-user")
    public String saveUser(@RequestBody User user){
        var a=this.service.saveUser(user);
        return a;
    }
    @GetMapping("/getuser-name/{name}")
    public List<User>getUserByName(@PathVariable String name){
        var u=this.service.getUserFromName(name);
        return u;
    }
    @GetMapping("/getuser-age/{age}")
    public List<User>getUserByAge(@PathVariable int age){
        var u=this.service.getUserFromAge(age);
        return u;
    }
    @GetMapping("/getuser-totalwallet/{total}")
    public List<User>getUserByTotalWallet(@PathVariable int total){
        var u=this.service.getUserFromTotalWallet(total);
        return u;

    }
    @GetMapping("/increment-userwallet/{email}")
    public String increaseTotal(@PathVariable String email){
        this.service.incrementWallet(email);
        return "increment success";
    }

}
