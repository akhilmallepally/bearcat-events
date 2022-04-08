from django.db import models

# Create your models here.
class EventModel(models.Model):

    title = models.CharField(max_length = 100, help_text='title of event')
    description = models.TextField(help_text='description about event')
    location = models.TextField(help_text='location of the event', blank=True)
    event_date = models.DateTimeField()
    image = models.ImageField(upload_to='images', blank=True)

    
    def __str__(self):
        return self.title

