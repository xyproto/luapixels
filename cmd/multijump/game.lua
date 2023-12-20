window_title = "Multi Jump"

local player_x = 100
local player_y = 100

local player_ax = 0
local player_ay = 4

local earth_ax = 0
local earth_ay = 1

local up_pressed = false
local down_pressed = false
local left_pressed = false
local right_pressed = false
local space_pressed = false

function at_start()
  print(window_title)
end

function at_every_frame()
  background(0, 128, 255)
  for y = 0, 199 do
    for x = 0, 255 do
      plot(x, y, x)
    end
    for x = 256, 318 do
      plot(x, y, 50)
    end
    plot(319, y, 70)
  end
  plot(318, 198, 41)
  plot(318, 199, 41)
  plot(319, 198, 41)
  plot(319, 199, 41)
  plot(player_x, player_y, 41)
end

function at_every_tick()
  -- Apply friction if player is on the ground and neither left nor right arrow key is held down
  if player_y == 199 and not left_pressed and not right_pressed then
    player_ax = player_ax * 0.7
  end

  -- Update player position based on key presses
  if left_pressed then
    player_ax = player_ax - 5
    player_x = player_x - 1
  end
  if right_pressed then
    player_ax = player_ax + 5
    player_x = player_x + 1
  end
  if space_pressed then
    player_ay = -10
    space_pressed = false  -- Prevent sustained jumping
  end

  -- Apply gravity
  player_ax = player_ax + earth_ax
  player_ay = player_ay + earth_ay

  -- Update player position
  player_x = player_x + player_ax
  player_y = player_y + player_ay

  -- Boundary checks
  if player_y > 199 then
    player_y = 199
    player_ay = -player_ay * 0.5
  elseif player_y < 0 then
    player_y = 0
    player_ay = -player_ay * 0.5
  end
  if player_x > 319 then
    player_x = 319
    player_ax = -player_ax * 0.5
  elseif player_x < 0 then
    player_x = 0
    player_ax = -player_ax * 0.5
  end
end

function at_left_pressed()
  left_pressed = true
end

function at_left_released()
  left_pressed = false
end

function at_right_pressed()
  right_pressed = true
end

function at_right_released()
  right_pressed = false
end

function at_space_pressed()
  space_pressed = true
end

function at_space_released()
  space_pressed = false
end

function at_esc_pressed()
  quit()
end

function at_key_pressed()
  -- was 'q' pressed?
  if last_key == 81 then
    quit()
  end
end

function at_end()
  print("bye!")
end
