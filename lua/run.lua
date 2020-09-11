function testGoFunc()
    local params = { GO_Util.printMe("来自lua的问候", "lehao ") }
    for k, v in pairs(params) do
        print(k, v)
    end
end

function initParams()
    local example = require('lua.db_module.example')
    local name = getParams("name")
    print(name)
    local code, tables = example.getList()

    return { code = code, response = tables }
end
