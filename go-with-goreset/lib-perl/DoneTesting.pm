use Task;
use Test::More;

package DoneTesting;
use Mo qw'build default builder coerce is required';
extends 'Task';
sub run{ 
	Test::More::done_testing(); 
}


1;
