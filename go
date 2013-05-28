    -- Tells the turtle to move in a direction, and maintains position information
    -- Copyright (C) 2013  Nathanael Cunningham

    -- This program is free software: you can redistribute it and/or modify
    -- it under the terms of the GNU General Public License as published by
    -- the Free Software Foundation, either version 3 of the License, or
    -- (at your option) any later version.

    -- This program is distributed in the hope that it will be useful,
    -- but WITHOUT ANY WARRANTY; without even the implied warranty of
    -- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    -- GNU General Public License for more details.

    -- You should have received a copy of the GNU General Public License
    -- along with this program.  If not, see <http://www.gnu.org/licenses/>.



args = {...}

if args[1] and args[1] == "help" then
  print("Usage: go {direction} numbers...\n"..
	"Sends the turtle in 'direction' num distance\n"..
	"goto is also available as a 'direction' with a 'x y z' coordinate\n"..
	"right and left are also available but do not take numbers")
  		
  return
end

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
