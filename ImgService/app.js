const express=require('express');
const {errorNotCaught,errorNotFound, authenticated}=require('./middlewares');
const cors=require('cors');
const upload=require('./upload')

const app=express();

app.use(express.json());
app.use(express.urlencoded({extended:false}));

app.use(cors({origin:true,credentials:true}));

app.use('/Upload',upload);
app.use("/Serve",express.static('resources/img'));
// app.get("/AuthTest",authenticated,(req,res)=>res.json({"messsage":"Auth success"}));

app.use(errorNotFound);
app.use(errorNotCaught);

module.exports=app;
