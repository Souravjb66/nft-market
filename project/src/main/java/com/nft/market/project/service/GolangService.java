package com.nft.market.project.service;

import com.nft.market.project.Togolang.WtihId;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GolangService{
    @Autowired
    private GoLangClient goClient;
    public String getName(String address){
        var name=goClient.GetTokenName(address);
        return name;
    }
    public String getSymbol(String address){
        var sym=goClient.GetTokenSymbol(address);
        return sym;
    }
    public String getTotalSupply(String address){
        var total=goClient.GetTotalSupply(address);
        return total;
    }
    public String getContractOwner(String address){
        var owner=goClient.GetOwnerOfTheContract(address);
        return owner;
    }
    public String getTokenValueById(WtihId with){
        var value=goClient.GetTokenValueById(with);
        return value;
    }
    public String getTokenAddressById(WtihId with){
        var address=goClient.GetTokenAddressById(with);

        return address;
    }
    public String getUserTotalToken(String address){
        var total=goClient.GetUserTotalToken(address);
        return total;
    }
    public String MintAToken(String name){
        var hash=goClient.MintMyToken(name);
        return hash;
    }
    public String ChangeTokenOwner(long id){
        var hash=goClient.ChangeTokenOwner(id);
        return hash;
    }


}
