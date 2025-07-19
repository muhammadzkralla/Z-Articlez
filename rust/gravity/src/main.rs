use once_cell::sync::Lazy;
use std::sync::Mutex;

const WIDTH: usize = 120;
const HEIGHT: usize = 36;
const GRAVITATIONAL_CONST: f64 = 1e-3;

static SCREEN: Lazy<Mutex<Vec<Vec<char>>>> = Lazy::new(|| {
    let screen = vec![vec![' '; WIDTH]; HEIGHT];
    Mutex::new(screen)
});

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

fn clear_screen() {
    // Acquire lock
    let mut screen = SCREEN.lock().unwrap();

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

fn update_pos(body1: &Body, body2: &mut Body) {
    // Acquire lock
    let mut screen = SCREEN.lock().unwrap();

    let vx = calc_vx(&body1, &body2);
    let vy = calc_vy(&body1, &body2);
    body2.vx += vx;
    body2.vy += vy;

    body2.update();

    let x2 = body2.x as isize;
    let y2 = body2.y as isize;
    let x1 = body1.x as isize;
    let y1 = body1.y as isize;

    if x2 >= 0 && x2 < WIDTH as isize && y2 >= 0 && y2 < HEIGHT as isize {
        screen[y2 as usize][x2 as usize] = body2.ch;
    }

    if x1 >= 0 && x1 < WIDTH as isize && y1 >= 0 && y1 < HEIGHT as isize {
        screen[y1 as usize][x1 as usize] = body1.ch;
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

fn render(body1: &Body, body2: &mut Body) {
    loop {
        clear_screen();
        update_pos(&body1, body2);

        // Clear terminal
        print!("\x1B[2J\x1B[H");

        let screen = SCREEN.lock().unwrap();
        for row in screen.iter() {
            let mut line = String::new();
            for ch in row.iter() {
                line.push(*ch);
            }

            println!("{}", line);
        }

        std::thread::sleep(std::time::Duration::from_millis(120));
    }
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
    let sun = Body::new(
        (WIDTH / 2) as f64,
        (HEIGHT / 2) as f64,
        0.0,
        0.0,
        '●',
        100000.0,
    );

    let mut body = Body::new(
        (WIDTH / 2) as f64 + 20.0,
        (HEIGHT / 2) as f64,
        0.0,
        1.8,
        '●',
        10.0,
    );

    render(&sun, &mut body);
}
