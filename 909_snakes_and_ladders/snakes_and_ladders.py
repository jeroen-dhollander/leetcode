import queue


# Huge number of steps used to indicate a square is unreachable
_UNREACHABLE = 1000000

FoundIt = Exception


class SnakesAndLadders(object):
    '''
        https://leetcode.com/problems/snakes-and-ladders/description/

        Given a snakes-and-ladders bord of dimensions NxN,
        find the minimum number of moves to reach the end.
        In each move you can move 1 to 6 squares forward.
    '''

    def __init__(self, board):
        self.board = board
        self.linear_board = list(_create_linear_board(board))
        # Contains the paths that we still must investigate
        self.paths_to_investigate = queue.Queue()

        self.shortest_path_to_end = None
        self.visited = [False] * self.num_squares

    def solve(self):
        try:
            start = Path(1, steps_to_get_here=0)
            self._do_next_move(start)

            while not self.paths_to_investigate.empty():
                path = self.paths_to_investigate.get(block=False)
                self._do_next_move(path)
        except FoundIt:
            pass

    @property
    def num_steps_to_end(self):
        return self.shortest_path_to_end.steps_to_get_here if self.shortest_path_to_end else -1

    def _get_shortest_path(self, index):
        return self.shortest_paths[index]

    def _do_next_move(self, path):
        for delta in range(1, 7):
            next_index = path.index + delta
            if next_index <= self.num_squares:
                self._evaluate_new_path(path.move_to(next_index))

    def _evaluate_new_path(self, new_path):
        # Follow ladder/snake if there is any
        if self._has_ladder_or_snake(new_path):
            new_path = new_path.ladder_to(self.square_value(new_path.index))

        if new_path.index == self.num_squares:
            self.shortest_path_to_end = new_path
            raise FoundIt()

        if self.visited[new_path.index]:
            return

        self.visited[new_path.index] = True
        self.paths_to_investigate.put(new_path)

    def _has_ladder_or_snake(self, path):
        return self.square_value(path.index) != -1

    @property
    def num_squares(self):
        return len(self.linear_board)

    def square_value(self, index):
        return self.linear_board[index - 1]


def _create_linear_board(board):
    height = len(board)
    width = len(board[0])

    reversed_ = False
    for row in range(height - 1, -1, -1):
        r = range(width - 1, -1, -1) if reversed_ else range(0, width, 1)
        for column in r:
            yield board[row][column]
        reversed_ = not reversed_


class Path(object):

    def __init__(self, index, previous_path=None, steps_to_get_here=_UNREACHABLE):
        self.index = index
        self.previous = previous_path
        self.previous_index = previous_path.index if previous_path else -1
        self.steps_to_get_here = steps_to_get_here

    @property
    def is_unreachable(self):
        return self.steps_to_get_here is _UNREACHABLE

    @property
    def is_start(self):
        return self.index == 1

    @property
    def movement(self):
        if self.is_start:
            return 'start'
        if self.previous.steps_to_get_here != self.steps_to_get_here:
            return 'moved'
        if self.previous.index < self.index:
            return 'ladder'
        if self.previous.index > self.index:
            return 'snake'
        return 'INVALID MOVEMENT'

    def move_to(self, next_square):
        return Path(
            index=next_square,
            previous_path=self,
            steps_to_get_here=self.steps_to_get_here + 1,
        )

    def ladder_to(self, next_square):
        return Path(
            index=next_square,
            previous_path=self,
            steps_to_get_here=self.steps_to_get_here,
        )

    def __str__(self):
        if self.is_unreachable:
            return '{}: unreachable'.format(self.index)
        elif self.is_start:
            return '1 (start)'
        else:
            return '{number} ({movement}) <- {previous}'.format(
                number=self.index,
                movement=self.movement,
                previous=str(self.previous),
            )

    def __repr__(self):
        return 'Path({}, {})'.format(self.index, self.movement)


class Solution:
    def snakesAndLadders(self, board):
        """
        :type board: List[List[int]]
        :rtype: int
        """
        board = SnakesAndLadders(board)
        board.solve()
        return board.num_steps_to_end


