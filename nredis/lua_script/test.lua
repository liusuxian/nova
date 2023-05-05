local result = redis.call('SET', KEYS[1], ARGV[1])
if not result then
  return result
end
result = redis.call('GET', KEYS[1])
return result
