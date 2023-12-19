window_title = "Simple Example"

function at_start()
    print("hi")
end

function at_every_frame()
    -- Set the background color to blue (RGB 0, 128, 255)
    background(0, 128, 255)
    -- At (100, 100), draw a red pixel (index 41 in the VGA palette)
    plot(100, 100, 41)
end

function at_keypress()
    quit()
end

function at_end()
    print("bye!")
end
