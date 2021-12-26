
-- http://localhost:8081/product/sku111210
--wrk -t1 -c1 -d5s -s ./testcase/wrkscript.lua http://localhost:8081
--wrk -t2 -c10 -d5s -s ./testcase/wrkscript.lua http://localhost:8081

request = function()
    wrk.headers["Connection"] = "Keep-Alive"
    skuid = math.random(1,1000)
    path = "/product/sku" .. skuid
    return wrk.format("GET", path)
  end