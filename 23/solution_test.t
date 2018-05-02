use strict;
use Test;

BEGIN { plan tests => 1 }

use rlib '.';
use Solution;

sub NewMatrix {
    my $size = shift;

    my @matrix = ();

    for (my $i=0; $i < $size; $i++) {
        for (my $j = 0; $j < $size; $j++) {
            $matrix[$i][$j] = 0;
        }
    }

    return @matrix;
}

sub SolutionTest {
    
    my @matrix = NewMatrix(4);
    
    $matrix[1][0] = 1;
    $matrix[1][1] = 1;
    $matrix[1][3] = 1;

    my $solution = new Solution(@matrix);

    my %start = ( x=>3, y=>0 );

    my %end = ( x=>0, y=>0 );
    
    my $dist = $solution->solve(\%start, \%end);

    ok($dist == 7);
}


SolutionTest();

