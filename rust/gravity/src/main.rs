const WIDTH: usize = 120;
const HEIGHT: usize = 70;
const GRAVITATIONAL_CONST: f64 = 1e-3;

struct Body {
    mass: f64,
    ch: char,
    x: f64,
    y: f64,
    vx: f64,
    vy: f64,
}

impl Body {
    fn new(x: f64, y: f64, vx: f64, vy: f64, ch: char, mass: f64) -> Self {
        Body {
            mass,
            ch,
            x,
            y,
            vx,
            vy,
        }
    }

    fn update(&mut self) {
        self.x += self.vx;
        self.y += self.vy;
    }
}

fn clear_screen(screen: &mut Vec<Vec<char>>) {
    for i in 0..HEIGHT {
        for j in 0..WIDTH {
            if i == 0 || i == HEIGHT - 1 {
                screen[i][j] = '─';
            } else if j == 0 || j == WIDTH - 1 {
                screen[i][j] = '│';
            } else {
                screen[i][j] = ' ';
            }
        }
    }

    // Draw corners
    screen[0][0] = '┌';
    screen[0][WIDTH - 1] = '┐';
    screen[HEIGHT - 1][0] = '└';
    screen[HEIGHT - 1][WIDTH - 1] = '┘';
}

fn update_pos(sun: &Body, planets: &mut Vec<Body>, screen: &mut Vec<Vec<char>>) {
    for planet in planets.iter_mut() {
        let vx = calc_vx(&sun, &planet);
        let vy = calc_vy(&sun, &planet);
        planet.vx += vx;
        planet.vy += vy;

        planet.update();

        let x2 = planet.x as isize;
        let y2 = planet.y as isize;
        let x1 = sun.x as isize;
        let y1 = sun.y as isize;

        if x2 >= 0 && x2 < WIDTH as isize && y2 >= 0 && y2 < HEIGHT as isize {
            screen[y2 as usize][x2 as usize] = planet.ch;
        }

        if x1 >= 0 && x1 < WIDTH as isize && y1 >= 0 && y1 < HEIGHT as isize {
            screen[y1 as usize][x1 as usize] = sun.ch;
        }
    }
}

fn render(screen: &mut Vec<Vec<char>>) {
    // Clear terminal
    print!("\x1B[2J\x1B[H");

    for row in screen.iter() {
        let mut line = String::new();
        for ch in row.iter() {
            line.push(*ch);
        }

        println!("{}", line);
    }
}

fn calc_vy(body1: &Body, body2: &Body) -> f64 {
    let fy = calc_fy(&body1, &body2);

    return fy / body2.mass;
}

fn calc_vx(body1: &Body, body2: &Body) -> f64 {
    let fx = calc_fx(&body1, &body2);

    return fx / body2.mass;
}

fn calc_fy(body1: &Body, body2: &Body) -> f64 {
    let force = calc_force(&body1, &body2);
    let dist = calc_dist(&body1, &body2);
    let dy = body1.y - body2.y;

    return force * (dy / dist);
}

fn calc_fx(body1: &Body, body2: &Body) -> f64 {
    let force = calc_force(&body1, &body2);
    let dist = calc_dist(&body1, &body2);
    let dx = body1.x - body2.x;

    return force * (dx / dist);
}

fn calc_force(body1: &Body, body2: &Body) -> f64 {
    let dist = calc_dist(&body1, &body2);
    return GRAVITATIONAL_CONST * (body1.mass * body2.mass) / (dist * dist);
}

fn calc_dist(body1: &Body, body2: &Body) -> f64 {
    let x1 = body1.x;
    let x2 = body2.x;
    let y1 = body1.y;
    let y2 = body2.y;

    let dx = x2 - x1;
    let dy = y2 - y1;
    let dist = (dx * dx) + (dy * dy);

    return dist.sqrt();
}

fn main() {
    let mut screen = vec![vec![' '; WIDTH]; HEIGHT];
    let mut planets = vec![
        // Mercury
        Body::new(
            (WIDTH / 2) as f64 + 10.0,
            (HEIGHT / 2) as f64,
            0.0,
            2.4,
            'M',
            2.0,
        ),
        // Venus
        Body::new(
            (WIDTH / 2) as f64 + 15.0,
            (HEIGHT / 2) as f64,
            0.0,
            2.0,
            'V',
            5.0,
        ),
        // Earth
        Body::new(
            (WIDTH / 2) as f64 + 20.0,
            (HEIGHT / 2) as f64,
            0.0,
            1.8,
            '🌍',
            10.0,
        ),
        // Mars
        Body::new(
            (WIDTH / 2) as f64 + 25.0,
            (HEIGHT / 2) as f64,
            0.0,
            1.6,
            'M',
            8.0,
        ),
        // Jupiter
        Body::new(
            (WIDTH / 2) as f64 + 35.0,
            (HEIGHT / 2) as f64,
            0.0,
            1.2,
            'J',
            100.0,
        ),
        // Saturn
        Body::new(
            (WIDTH / 2) as f64 + 45.0,
            (HEIGHT / 2) as f64,
            0.0,
            1.0,
            'S',
            80.0,
        ),
        // Uranus
        Body::new(
            (WIDTH / 2) as f64 + 55.0,
            (HEIGHT / 2) as f64,
            0.0,
            0.8,
            'U',
            60.0,
        ),
        // Neptune
        Body::new(
            (WIDTH / 2) as f64 + 65.0,
            (HEIGHT / 2) as f64,
            0.0,
            0.7,
            'N',
            60.0,
        ),
    ];

    let sun = Body::new(
        (WIDTH / 2) as f64,
        (HEIGHT / 2) as f64,
        0.0,
        0.0,
        '●',
        100000.0,
    );

    loop {
        clear_screen(&mut screen);
        update_pos(&sun, &mut planets, &mut screen);

        render(&mut screen);

        std::thread::sleep(std::time::Duration::from_millis(120));
    }
}
