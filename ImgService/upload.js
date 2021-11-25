const express = require('express');
const router = express.Router();
const {uploadImage,authenticated}=require('./middlewares');

// router.get("/AuthTest",authenticated,(req,res)=>res.json({"messsage":"Auth success"}));

router.post('/',authenticated,uploadImage().single('img'),async (req,res) => {
    return res.json({"message":"Upload successful"});
});



module.exports=router
