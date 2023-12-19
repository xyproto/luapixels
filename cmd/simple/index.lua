window_title = "Pixel McPixelface"

function at_start()
    print("hi")
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
    plot(319, 199, 41)
    plot(319, 199, 41)
    plot(319, 199, 41)
    plot(319, 199, 41)
end

function at_end()
    print("bye!")
end
