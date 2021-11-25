const multer=require('multer');
const path=require('path');
const http=require('http');
const got=require('got');

const authenticated=async function (req,res,next){
    const authHeader=req.headers.authorization;
    if(!authHeader){
        return res.status(403).json({"message":"Unauthorized"});
    }

    try {
        const response=await got.get(process.env.AUTH_SERVICE,{headers:{"Authorization":authHeader},timeout:{response:2000}});
        if (!JSON.parse(response.body).hasOwnProperty("user") || response.statusCode!==200){
            console.log('Authentication failed!');
            return res.status(403).json({"message":"Unauthorized"});
        }
        console.log("Authentication successful!");
        next();
    } catch (e) {
        return res.status(500).json({"message":e.message});
    }
}



const errorNotFound=(req,res,next)=>{
    const error=new Error("Not found")
    error.status=404
    next(error)
}
const errorNotCaught=(error,req,res,next)=> {
    res.status(error.status ?? 500)
    if(process.env.NODE_ENV==='development')
        return res.json({
            message:error.message,
            stack:error.stack
        })
    return res.json({
        message:error.message
    })
}


const uploadImage = ()=>{
    const storage=multer.diskStorage({
        destination:(req,file,cb)=>{cb(null,path.resolve('resources/img'))},
        filename:(req,file,cb)=>{cb(null,file.originalname)}
    })
    return multer({storage:storage});
}

module.exports = {errorNotFound,errorNotCaught,authenticated,uploadImage}
