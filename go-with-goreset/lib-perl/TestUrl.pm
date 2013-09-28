use HttpTask;

package TestUrl;
use Mo qw'build default builder coerce is required';
extends 'HttpTask';
has tests => (default => sub{[]});
sub run
{
	my $self = shift;
	my $content = GET->new(url => $self->url )->run;
	for(@{ $self->tests} )
	{
		$_->($content);
	}
}

1;
