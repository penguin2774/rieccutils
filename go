args = {...}

if #args <1 then
  print("usage: go {direction} nums...")
  return
end

function numorone(x)
  if not x then
    return(1)
  else
    return(tonumber(x))
  end
end

if args[1] == "forward" then
  move.forward(numorone(args[2]))
  return
elseif args[1] == "left" then
  move.turnLeft()
  return
elseif args[1] == "right" then
  move.turnRight()
  return
elseif args[1] == "up" then
  move.up(numorone(args[2]))
  return
elseif args[1] == "down" then
  move.down(numorone(args[2]))
  return
elseif args[1] == "goto" then
  move.goto(tonumber(args[2]),
            tonumber(args[3]),
            tonumber(args[4]))
  return
elseif args[1] == "north" then
  move.north(numorone(args[2]))
  return
elseif args[1] == "south" then
  move.south(numorone(args[2]))
  return
elseif args[1] == "east" then
  move.east(numorone(args[2]))
  return
elseif args[1] == "west" then
  move.west(numorone(args[2]))
  return
elseif args[1] == "home" then
  move.goto(0,0,0)
  move.face(move.NORTH)
else
  print("I don't know how to ", args[1])
  return
end
