# MazeSolver enables DockerMaze LEGS robot module resolve ASCII Art mazes.
# Unfortunately, LEGS robot module is damaged and is not possible assemble
# it successfully.
# Repair MazeSolver library in order to provide high performant DockerMaze
# LEGS module that can be assaembled to your robot.

init_coord = (1, 0)

# Directions: 1 - left, 2 - up, 3 - left, 4 - down
directions = [4, 1, 2, 3, 4, 1, 2, 3, 4]
direction_letters = {1: 'LEFT', 2: 'UP', 3: 'LEFT', 4: 'DOWN'}


class MazeSolver:
    def __init__(self, maze):
        self.last_move_direction = 2
        self.route = []
        self.maze = maze.split('\n')

    def calculate_end_coord(self):
        width = len(self.maze[0])
        height = len(self.maze)
        return (height-1-1, width-1)

    def move(self, coord, direction):
        if direction == 1:
            return (coord[0], coord[1]+1)
        elif direction == 2:
            return (coord[0]-1, coord[1])
        elif direction == 3:
            return (coord[0], coord[1]-1)
        elif direction == 4:
            return (coord[0]+1, coord[1])

    def is_valid_coord(self, coord):
        if self.maze[coord[0]][coord[1]] == ' ':
            return True
        else:
            return False

    def make_move(self, coord):
        for i in range(self.last_move_direction-1, self.last_move_direction+3):
            new_coord = self.move(coord, directions[i])
            if self.is_valid_coord(new_coord):
                self.last_move_direction = directions[i]
                self.route.append(direction_letters[directions[i]])
                return new_coord

    def solve_maze(self):
        end_coord = self.calculate_end_coord()
        coord = init_coord
        while coord != end_coord:
            coord = self.make_move(coord)
        return " ".join(self.route)
