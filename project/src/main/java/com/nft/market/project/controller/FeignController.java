package com.nft.market.project.controller;

import com.nft.market.project.Togolang.WtihId;
import com.nft.market.project.service.GolangService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class FeignController {
    @Autowired
    private GolangService service;
    @PostMapping("/token-name")
    public String getName(@RequestBody String address){
        var a=service.getName(address);
        return a;
    }
    @PostMapping("/token-symbol")
    public String getSymbol(@RequestBody String address){
        var sym=service.getSymbol(address);
        return sym;

    }
    @PostMapping("/total-supply")
    public String getTotalSupply(@RequestBody String address){
        var total=service.getTotalSupply(address);
        return total;
    }
    @PostMapping("/user-total-token")
    public String getTotalUserToken(@RequestBody String address){
        var total=service.getUserTotalToken(address);
        return total;
    }
    @PostMapping("/mint-my-token")
    public String MintMyToken(@RequestBody String name){
        var hash=service.MintAToken(name);
        return hash;
    }
    @PostMapping("/change-token-owner")
    public String ChangeTokenOwner(@RequestBody long id){
        var hash=service.ChangeTokenOwner(id);
        return hash;
    }
    @PostMapping("/token-value-byid")
    public String getTokenValueById(@RequestBody WtihId with){
        var value=service.getTokenValueById(with);
        return value;
    }
    @PostMapping("/token-address-byid")
    public String getTokenAddressById(@RequestBody WtihId with){
        var address=service.getTokenAddressById(with);
        return address;
    }

}
