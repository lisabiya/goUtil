---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by wakfu.
--- DateTime: 2020/9/7 9:32
---


--直接引用声明模块
function TestHttp(params)
    local code, response = httprequest.get(
            { url = "https://www.wanandroid.com/hotkey/json" })
    --调用函数
    return { code = code, response = response, params = params }
end

