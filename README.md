# go-sciio — brand assets

Official logos for the **go-sciio** organization.

<img src="avatar/go-sciio.png" width="160" alt="go-sciio">

The mark is what I/O means: data going out along one line and coming back along
another, two opposite arrows. It is written by `mark/main.go`, which derives
every coordinate from the length of the shafts, the size of the heads and the
distance between the lines, and centres the drawing; nothing is placed by hand.
Its gradient (`#8FA3BF` to `#3E516B`) is a slate that sets go-sciio apart from
the violet of go-numerics, the emerald of go-ndarray and go-fft, the blue of
go-plot, the magenta of go-bignum, the coral of go-sparse, the orange of
go-approx and the rose of go-dsp; its dark end reads on white at 8.10:1 and its
light end on a dark page at 7.51:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-sciio.svg` differs.

---
*Generated from `svg/color/go-sciio.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
