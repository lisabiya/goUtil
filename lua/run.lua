require("lua/tools")

local params = { GO_Util.printMe("来自lua的问候") }

for k, v in pairs(params) do
    print(k,v)
end