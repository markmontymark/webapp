package Task;
use Mo;

sub run 
{
	my $self = shift;
	die "Task::run not overridden in subclass, $self, with args ", @_,"\n";
}

1;
