package com.nft.market.project.repository;

import com.nft.market.project.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface UserRepository extends MongoRepository<User,String> {
    List<User>findByname(String name);
    List<User>findByemail(String email);
    List<User>findByage(int age);
    List<User>findBytotalWallet(int no);



}
