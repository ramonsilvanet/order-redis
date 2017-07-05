function postMerchant()
    wrk.headers["merchant-id"] = "1234"
    return wrk.format("POST", "/merchant")
end

function getMerchant()
    return wrk.format("GET", "/merchant/1234")
end

init = function(args) 

    local r = {}
    r[1] = postMerchant()
    r[2] = getMerchant()

    req = table.concat(r)
end

request = function()
    return req
end