M = {}

-- Stream handle API of Lua HTTP filter
-- @doc https://www.envoyproxy.io/docs/envoy/v1.7.0/configuration/http_filters/lua_filter#config-http-filters-lua-stream-handle-api
function M.log(handle)
    handle:logDebug("[mylibary.M.log] started")

    local index = 0
    for chunk in handle:bodyChunks() do
        handle:logDebug('[mylibrary.M.log] showing bodyChunks')

        local len = chunk:length()
        local result = chunk:getBytes(index, len)
        index = index + len

        handle:logDebug(result)

        -- Make an HTTP call.
        handle:logDebug("[mylibary.M.log] sending log to `backend`")

        local headers, body = handle:httpCall(
         "log-service",
            {
              [":method"] = "POST",
              [":path"] = "/log",
              [":authority"] = "lua_cluster"
            },
          result,
          24224)

        handle:logDebug("[mylibary.M.log] sent log to `backend`")
    end

    handle:logDebug("[mylibary.M.log] finished")
end

return M
