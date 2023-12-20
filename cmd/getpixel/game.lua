window_title = "Get Pixel"

function at_every_frame()
  -- At (100, 100), draw a pixel using color 7 (the index in the VGA palette)
  plot(100, 100, 7)
end

function at_key_pressed()
  quit()
end

function at_end()
  print("pixel at (100, 100) has palette index " .. getpixel(100, 100))
end
