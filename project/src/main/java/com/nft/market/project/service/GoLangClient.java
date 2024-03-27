package com.nft.market.project.service;

import com.nft.market.project.Togolang.WtihId;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "golang-client", url = "${golang.baseurl}")
public interface GoLangClient {

    @PostMapping("/token-name") // Replace "/endpoint" with the actual endpoint in your GoLang application
    String GetTokenName(@RequestBody String address);
    @PostMapping("/token-symbol")
    String GetTokenSymbol(@RequestBody String address);
    @PostMapping("/total-supply")
    String GetTotalSupply(@RequestBody String address);
    @PostMapping("/owner-of-the-contract")
    String GetOwnerOfTheContract(@RequestBody String address);
    @PostMapping("/token-value-by-id")
    String GetTokenValueById(@RequestBody WtihId with);
    @PostMapping("/token-address-by-id")
    String GetTokenAddressById(@RequestBody WtihId with);
    @PostMapping("/mint-token")
    String MintMyToken(@RequestBody String name);
    @PostMapping("/change-token-owner")
    String ChangeTokenOwner(@RequestBody long id);
    @PostMapping("/user-total-token")
    String GetUserTotalToken(@RequestBody String address);

}
