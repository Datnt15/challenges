package Solution;

use Data::Dumper;

sub new {
    my $class = shift;
    my $self = {
        _input => shift,
        _visited => (),
    };
    bless $self, $class;
    return $self;
}

# static increments for directions
my @directions = ( {x => -1, y => 0}, 
                 {x => 1, y => 0},
                 {x => 0, y => -1},
                 {x => 0, y => 1} );


sub isValid {
    my ($self, $col, $row) = @_;

    my @input = @{ $self->{_input} };
    my @visited = @{ $self->{_visited} };

    if ($input[$col][$row] || $visited[$col][$row]) {
        print "$col, $row not valid\n";
        return 0;
    }

    my $size = scalar @input;

    return $col >= 0 && $row >= 0 && $row < $size && $col < $size;
}

sub visit {
    my ($self, $col, $row) = @_;

    $self->{_visited}[$col][$row] = 1;
}

sub solve {
    my ($self, $start, $end) = @_;

    $start->{dist} = 1;

    my @queue = ($start);
    
    print Dumper($start);
    print Dumper($end);

    $self->visit($start->{x}, $start->{y});

    while (@queue) {

        my $point = shift @queue;

        # test if we found the end position
        if ($point->{x} == $end{x} && $point->{y} == $end{y}) {
            return $point->{dist};
        }

        # scan each direction
        for my $direction (@directions) {

            my $col = $point->{x} + $direction->{x};
            my $row = $point->{y} + $direction->{y};

            if ($self->isValid($col, $row)) {
                $self->visit($col, $row);
                my %next = ( x => $col, y => $row, dist => $point->{dist} + 1 );
                push @queue, \%next;
            }
        }
    }

    # did not find the end position
    return -1;
}

1;

