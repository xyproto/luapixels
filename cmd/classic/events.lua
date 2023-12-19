window_title = "Pixel McPixelface"

player_x = 100
player_y = 100

player_ax = 0
player_ay = 0

earth_ax = 0
earth_ay = 1

max_run_speed = 5
run_acceleration = 0.5
friction = 0.85
jump_strength = 10
gravity = 0.5

is_on_ground = false
up_pressed = false
down_pressed = false
left_pressed = false
right_pressed = false
space_pressed = false

function at_start()
    print("Classic")
end

function at_every_tick()
    -- Apply friction if on the ground
    if is_on_ground then
        player_ax = player_ax * friction
    end

    -- Apply gravity
    player_ay = player_ay + gravity

    -- Horizontal movement
    if left_pressed then
        player_ax = player_ax - run_acceleration
        if player_ax < -max_run_speed then
            player_ax = -max_run_speed
        end
    elseif right_pressed then
        player_ax = player_ax + run_acceleration
        if player_ax > max_run_speed then
            player_ax = max_run_speed
        end
    end

    -- Jumping
    if space_pressed and is_on_ground then
        player_ay = -jump_strength
        is_on_ground = false
        space_pressed = false  -- Prevent multi-jumping
    end

    -- Update player position
    player_x = player_x + player_ax
    player_y = player_y + player_ay

    -- Collision with ground
    if player_y >= 199 then
        player_y = 199
        player_ay = 0
        is_on_ground = true
    end

    -- Collision with screen boundaries
    if player_x > 319 then
        player_x = 319
        player_ax = 0
    elseif player_x < 0 then
        player_x = 0
        player_ax = 0
    end
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

function at_keypress()
    local key = last_key
    -- Quit if GLFW key codes 81 ('q') or 256 ('Esc') are pressed
    if key == 81 or key == 256 then
        quit()
    end
    -- Adjust cursor position based on arrow keys or WASD
    -- GLFW key codes: 87 ('W'), 65 ('A'), 83 ('S'), 68 ('D')
    -- Arrow keys: 262 (Right), 263 (Left), 264 (Down), 265 (Up)
    if key == 87 or key == 265 then
        up_pressed = true
    elseif key == 83 or key == 264 then
        down_pressed = true
    elseif key == 65 or key == 263 then
        left_pressed = true
    elseif key == 68 or key == 262 then
        right_pressed = true
    elseif key == 32 then
        space_pressed = true
    end

end

function at_keyrelease()
    local key = last_key
    -- Adjust cursor position based on arrow keys or WASD
    -- GLFW key codes: 87 ('W'), 65 ('A'), 83 ('S'), 68 ('D')
    -- Arrow keys: 262 (Right), 263 (Left), 264 (Down), 265 (Up)
    if key == 87 or key == 265 then
        up_pressed = false
    elseif key == 83 or key == 264 then
        down_pressed = false
    elseif key == 65 or key == 263 then
        left_pressed = false
    elseif key == 68 or key == 262 then
        right_pressed = false
    elseif key == 32 then
        space_pressed = false
    end
end

function at_end()
    print("bye!")
end
