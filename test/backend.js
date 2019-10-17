const axios = require('axios')
const chai = require('chai')
const assert = chai.assert

require('dotenv').config()
const port = process.env.PORT
const client_id = process.env.client_id
const client_secret = process.env.client_secret
const validUserNames = []
const invalidUserNames = []

//before the test is done, let's randomly get a bunch of valid and invalid usernames.
before(async() => {
    //let's generate some random numbers
    randFirst = Math.floor(Math.random() * 20000) + 1  
    randSecond = randFirst + 1
    respFirst =  get(`https://api.github.com/users?since=${randFirst}&client_id=${client_id}&client_secret=${client_secret}`)
    respSecond = get(`https://api.github.com/users?since=${randSecond}&client_id=${client_id}&client_secret=${client_secret}`)
    await Promise.all([respFirst,respSecond]).then((response) => {
        userList = response[0].data.concat(response[1].data)
        userList.forEach((user) => {
            //We also generate same number of randomNames for every valid username
            let randomName = Math.random().toString(36).substring(1);
            validUserNames.push(user.login)
            invalidUserNames.push(randomName)
        })
        console.log("Generated VALID user IDs are:")
        console.log(validUserNames)    
        console.log("Generetad INVALID userIDs are:")
        console.log(invalidUserNames)
    })
})

describe('Test if the server is online',() => {
    it('should return a 200 Status code',async() => {
        resp = await post('http://localhost:9999/api/v1/getUser',{
            username:"syahrul12345"
        })
        assert.equal(resp.status,200,'NOT EQUAL')
    })
    console.log("end of test 1")
})

describe('Check if our Golang API can get userdata',() => {
    it('should say this',async() => {
        validUserNames.forEach((user) => {
            it(`should return a list when called to http://localhost:${port}/api/v1/getUser for ${user}`,async() => {
                resp = post('http://localhost:${port}/api/v1/getUser',{username:user})
                console.log(resp.data.result)
            })
        })
    })
    console.log("end of test")
})
function post(url,payload) {
    return axios.post(url,payload)
        .then((response) => {
            return response
        }).catch((error) => {
            return error
        })
}
function get(url) {
    return axios.get(url)
        .then((response) => {
            return response
        }).catch((error) => {
            return error
        })

}