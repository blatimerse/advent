#!/usr/bin/env perl

my $re = qr/(.*) bags contain(?: (\d+) (.*?) bags?[,.])+/;

while (<>){
    chomp;
    if (/$re/) {
        print "$1\n";
        print "$2\n";
        print "$3\n";
        print "$4\n";
        print "$5\n";
        next;
    }
    die "no match for $_\n"
}