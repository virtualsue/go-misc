#!/usr/bin/perl
use warnings;
use strict;

my $n = 5000000;
my $delta = 1/$n;
for ($i = 1; $i <= $n; $i++) {
    $x = ($i - 0.5) * $delta;
	$sum += 1.0 / (1.0 + $x*$x);
}
print 4 * $delta * $sum, "\n";
