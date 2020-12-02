#!/usr/bin/perl -w

# guidelines:
# leave numbers in string format as long as possible
# unpack and repack for operations.

my $digit_length = 4;

my $uno = "100";
my $dos = "42";
my $tres = "9999";

my $result;

#$result = &add( $uno, $dos );

my @sample;

$sample[2] = "1";
$sample[1] = "99";
$sample[0] = "321";

print &arb_pack( \@sample ) . "\n";

$test = "1099321";

print &arb_pack( &arb_unpack( $test ) ) . "\n";

$long_test = "10000000000000000000000000000000000000000000000000000000000000000000000000011111111111111111111111111111111122222222222222222222222222222333333333333333333333333333444444444444444444";

print &arb_pack( &arb_unpack( $long_test ) ) . "\n";

print "1 + 2 = " . &add( "1", "2" ) . qq(\n);;

print "1000 + 9000 = " . &add( "1000", "9000" ) . "\n";

exit( 0 );

sub add {
	my( @terms ) = @_;

	my( $result ) = "0";

	#this thing is still off by one...

	foreach $term ( @terms ) {
		print qq(adding $term to $result\n);
		my( $unpacked_result ) = &arb_unpack( $result );
		my( $unpacked_term ) = &arb_unpack( $term );

		my( $overflow ) = 0;

		for( $i = 0; $i < scalar( @$unpacked_term ); $i += 1 ) {
			print qq(i: ) . $i . qq( t: ) . @$unpacked_term[$i] . qq(\n);
			if( ! defined( @$unpacked_result[$i] ) ) {
				@$unpacked_result[$i] = @$unpacked_term[$i];
			} else {
				# do how should we deal with overflow
				print qq(adding: ) . @$unpacked_result[$i] . " to " . @$unpacked_term[$i] . "\n";
				@$unpacked_result[$i] += @$unpacked_term[$i];
				#$overflow = 0;
				if( length( qq(@$unpacked_result[$i]) ) > $digit_length ) {
					@$unpacked_result[$i+1] = "1";
					#$overflow = 1;
					print qq(ur: ) . @$unpacked_result[$i] . "\n";
					@$unpacked_result[$i] = @$unpacked_result[$i] =~ s/^.//;
					print qq(trimmed: ) . @$unpacked_result[$i] . "\n";
				}
				print join( ":", @$unpacked_result ) . "\n";
			}
			# if( $overflow ) {
			# 	@$unpacked_result[ scalar @$unpacked_result ] = "1";
			# 	$overflow = 0;
			# }
				print join( ":", @$unpacked_result ) . "\n";
		}

		$result = &arb_pack( $unpacked_result );
		print qq(result so far: $result\n);
	}

	return $result;
}



sub arb_unpack {
	my( $packed ) = @_;
 	my( @unpacked );

	# need to account for digit length.  Perhaps prepad baed on the length mod digit_size?

	if( length( $packed ) % $digit_length ) {
		$packed = "0" x ($digit_length - length($packed) % $digit_length ) . $packed;
	}

	@unpacked = reverse( unpack( "(A" . $digit_length . ")*", $packed ) );

	# foreach $d (@unpacked) {
	# 	print qq(what $d\n);
	# }

 	return \@unpacked;
}

sub arb_pack {
	my( $unpacked ) = @_;
	my( $packed );

	my( $first ) = 1;

	foreach $digit( reverse @$unpacked ) {
		if( $first ) {
			$packed = qq($digit);
			$first = 0;
		} else {
			my $padding = '';
			if( length $digit < $digit_length ) {
				$padding = "0" x ( $digit_length - length $digit );
			}
			$packed .= $padding . qq($digit);
		}
	}

	return $packed;
}

