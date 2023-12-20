window_title = "Simple Example"

function at_start()
    print("hi")
    -- Set color at palette index 7 to red (RGB 255, 0, 0)
    setpal(7, 255, 0, 0)
end

function at_every_frame()
    -- Set the background color to blue (RGB 0, 128, 255)
    background(0, 128, 255)
    -- At (100, 100), draw a red pixel (index 7 in the palette)
    plot(100, 100, 7)
end

function at_keypress()
    quit()
end

function at_end()
    print("bye!")
end
